import styles from '../styles/SideMenu.module.scss';
import Logo from './Logo';
import Navigation from './Navigation';
import User from './User';

export default function SideMenu(){
    return (<section className={styles.side_menu}>
        <div className={styles.menu_container}>
            <Logo />
            <Navigation />
        </div>
        <User name=''></User>
    </section>);
}