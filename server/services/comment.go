package services

import (
	"errors"
	"fmt"
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/repositories"
)

func CreateComment(trackID string, comment models.Comment) (models.Comment, error) {

	track, err := GetOneTrack(trackID)
	if err != nil {
		return models.Comment{}, err
	}

	comment.Track_id = trackID

	createdComment, err := repositories.CreateComment(comment)
	if err != nil {
		return models.Comment{}, err
	}

	track.Comments = append(track.Comments, createdComment.Id)
	_, err = UpdateTrack(track)
	if err != nil {
		fmt.Println(err)
	}
	return createdComment, nil
}

func GetComments(trackID string) ([]models.Comment, error) {

	track, err := GetOneTrack(trackID)
	if err != nil {
		return nil, err
	}
	var comments []models.Comment
	for _, commentID := range track.Comments {
		comment, err := repositories.GetComment(commentID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func DeleteComment(commentID string, trackID string) (models.Comment, error) {

	comment, err := repositories.GetComment(commentID)
	if err != nil {
		fmt.Println(comment)
		return models.Comment{}, err
	}

	if comment.Track_id != trackID {
		return models.Comment{}, errors.New("given Track id and Comment's Track id do not match")
	}

	track, err := GetOneTrack(trackID)
	if err != nil {
		return models.Comment{}, nil
	}

	foundFlag := false
	for i, commentIdx := range track.Comments {
		if commentIdx == commentID {
			foundFlag = true
			track.Comments = append(track.Comments[:i], track.Comments[i+1:]...)
			break
		}
	}

	if !foundFlag {
		return models.Comment{}, errors.New("comment not found")
	}

	_, err = repositories.UpdateTrack(track)
	if err != nil {
		return models.Comment{}, err
	}

	err = repositories.DeleteComment(commentID)
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func UpdateComment(comment models.Comment) (models.Comment, error) {

	_, err := GetOneTrack(comment.Track_id)
	if err != nil {
		return models.Comment{}, err
	}

	newComment, err := repositories.EditComment(comment)
	if err != nil {
		fmt.Println(err)
		return models.Comment{}, err
	}

	return newComment, nil
}
