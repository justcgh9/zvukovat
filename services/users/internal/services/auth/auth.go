package auth

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/justcgh9/zvukovat/services/users/internal/domain/models"
	"github.com/justcgh9/zvukovat/services/users/internal/lib/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
    log *slog.Logger
    usrSaver UserSaver
    usrProvider UserProvider
    usrUpdater UserUpdater
    tknSaver TokenSaver
    tknProvider TokenProvider
    tknRemover TokenRemover
    emlSender EmailSender
    accessTokenTTL time.Duration
    refreshTokenTTL time.Duration
    accessSecret string
    refreshSecret string
}

func New(
    logger *slog.Logger,
    usrSaver UserSaver,
    usrProvider UserProvider,
    usrUpdater UserUpdater,
    tknSaver TokenSaver,
    tknProvider TokenProvider,
    tknRemover TokenRemover,
    emlSender EmailSender,
    accessTokenTTL time.Duration,
    refreshTokenTTL time.Duration,
    accessSecret string,
    refreshSecret string,
) *Auth {
    return &Auth{
        log:             logger,
        usrSaver:        usrSaver,
        usrProvider:     usrProvider,
        usrUpdater:      usrUpdater,
        tknSaver:        tknSaver,
        tknProvider:     tknProvider,
        tknRemover:      tknRemover,
        emlSender:       emlSender,
        accessTokenTTL:  accessTokenTTL,
        refreshTokenTTL: refreshTokenTTL,
        accessSecret: accessSecret ,
        refreshSecret: refreshSecret,
    }
}

type UserSaver interface {
    SaveUser(
        ctx context.Context,
        usr models.User,
    ) (models.UserDTO, error)
}

type UserProvider interface {
    User(ctx context.Context, email string) (models.User, error)
}

type UserUpdater interface {
    UpdateFavourites(ctx context.Context, user models.UserDTO) error
    ActivateUser(ctx context.Context, link string) (models.UserDTO, error)
}

type TokenSaver interface {
    SaveToken(
        ctx context.Context,
        tkn models.Token,
    ) (models.Token, error)
}

type TokenProvider interface {
    Token(ctx context.Context, tknStr string) (models.Token, error)
}

type TokenRemover interface {
    DeleteToken(ctx context.Context, usrId string) error
}

type EmailSender interface {
    SendEmail(email, link string) error
}

func (a *Auth) SignUp(ctx context.Context, usr models.User, domainName string) (map[string]string, models.UserDTO, error) {
    const op = "services.auth.SignUp"
    var err error

    log := a.log.With(
        slog.String("op", op),
        slog.String("email", usr.Email),
        )

    log.Info("attempting to register user")

    checkUsr, _ := a.usrProvider.User(ctx, usr.Email)
    if checkUsr.Email == usr.Email  {
        log.Error("user with this email already exists")
        return nil, models.UserDTO{}, fmt.Errorf("%s: user with this email already exists", op)
    }

    usr.Password, err = hashPassword(usr.Password)
    if err != nil {
        log.Error(err.Error())
        return nil, models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }

    usr.ActivationLink = uuid.New().String()
    usrDTO, err := a.usrSaver.SaveUser(ctx, usr)
    if err != nil {
        log.Error(err.Error())
        return nil, models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }
    log.Info("user successfully created")

    err = a.emlSender.SendEmail(usr.Email, domainName + usr.ActivationLink)
    if err != nil {
        log.Error(err.Error())
        return nil, models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }
    log.Info("confirmation link sent", slog.String("link", usr.ActivationLink))

    tokens, err := jwt.GenerateTokens(usrDTO, a.accessSecret, a.refreshSecret)
    if err != nil {
        log.Error(err.Error())
        return nil, models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }

    token := models.NewToken(usrDTO.Id, tokens["refreshToken"])
    _, err = a.tknSaver.SaveToken(ctx, token)
    if err != nil {
        log.Error(err.Error())
        return nil, models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }

    return tokens, usrDTO, nil
}

func (a *Auth) SignIn(ctx context.Context, user models.User) (map[string]string, models.UserDTO, error) {
    const op = "services.auth.SignIn"

    log := a.log.With(
        slog.String("op", op),
        slog.String("email", user.Email),
        )

    log.Info("attempting to log in a user")

    usr, err := a.usrProvider.User(ctx, user.Email)
    if err != nil {
        log.Error(err.Error())
        return nil, models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }

    if !checkPasswordHash(user.Password, usr.Password) || !usr.IsActivated {
        log.Error(err.Error())
        return nil, models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }

    tkns, err := jwt.GenerateTokens(usr.UserDTO, a.accessSecret, a.refreshSecret)
    if err != nil {
        log.Error(err.Error())
        return nil, models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }

    tkn := models.NewToken(usr.Id, tkns["refreshToken"])

    _, err = a.tknSaver.SaveToken(ctx, tkn)
    if err != nil {
        log.Error(err.Error())
        return nil, models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }

    return tkns, usr.UserDTO, nil

}

func (a *Auth) SignOut(ctx context.Context, token string) error {
    const op = "services.auth.SignOut"

    log := a.log.With(
        slog.String("op", op),
        )

    claims, err := jwt.ValidateAccessToken(token, a.accessSecret)
    if err != nil {
        log.Error(err.Error())
        return fmt.Errorf("%s: %w", op, err)
    }

    err = a.tknRemover.DeleteToken(ctx, claims.Payload.Id)
    if err != nil {
        log.Error(err.Error())
        return fmt.Errorf("%s: %w", op, err)
    }

    return nil
}

func (a *Auth) ActivateUser(ctx context.Context, link string) (models.UserDTO, error) {
    const op = "services.auth.ActivateUser"

    log := a.log.With(
        slog.String("op", op),
        )

    usr, err := a.usrUpdater.ActivateUser(ctx, link)
    if err != nil {
        log.Error(err.Error())
        return models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }

    return usr, nil
}

func (a *Auth) Refresh(ctx context.Context, token string) (map[string]string, error) {
    const op = "services.auth.Refresh"

    log := a.log.With(
        slog.String("op", op),
        )

    log.Info("attempting to log out a user")
    claims, err := jwt.ValidateRefreshToken(token, a.refreshSecret)
    if err != nil {
        log.Error(err.Error())
        return nil, fmt.Errorf("%s: %w", op, err)
    }

    if !claims.Payload.UserDTO.IsActivated {
        log.Error("user is not activated")
        return nil, fmt.Errorf("%s: %s", op, "user is not activated")
    }

    _, err = a.tknProvider.Token(ctx, token)
    if err != nil {
        log.Error(err.Error())
        return nil, fmt.Errorf("%s: %w", op, err)
    }

    usr, err := a.usrProvider.User(ctx, claims.Payload.Email)
    if err != nil {
        log.Error(err.Error())
        return nil, fmt.Errorf("%s: %w", op, err)
    }

    tkns, err := jwt.GenerateTokens(usr.UserDTO, a.accessSecret, a.refreshSecret)
    if err != nil {
        log.Error(err.Error())
        return nil, fmt.Errorf("%s: %w", op, err)
    }

    tkn := models.NewToken(usr.Id, tkns["refreshToken"])
    _, err = a.tknSaver.SaveToken(ctx, tkn)
    if err != nil {
        log.Error(err.Error())
        return nil, fmt.Errorf("%s: %w", op, err)
    }

    return tkns, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
