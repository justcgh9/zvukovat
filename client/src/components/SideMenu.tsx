import styles from '../styles/SideMenu.module.scss';
import Logo from './Logo';

export default function SideMenu(){
    return (<section className={styles.side_menu}>
        <div className={styles.menu_container}>
            <Logo />
        </div>
        
        hehe side bar
    </section>);
}