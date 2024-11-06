import styles from '../styles/SideMenu.module.scss';
import Logo from './Logo';
import Navigation from './Navigation';
import NavLink from './NavLink';
import User from './User';
import CreateIcon from '../assets/create.svg'; // change to create (+ in soft square)

export default function SideMenu(){
    return (<section className={styles.side_menu}>
        <div className={styles.menu_container}>
            <Logo />
            <Navigation />
            <NavLink src={CreateIcon} text='Create track' to='/track/create'/>
        </div>
        <User name=''></User>
    </section>);
}


