import styles from "./page.module.css"

export default function Home() {
  return (
      <>
        <div className={styles.center}>
          <h1>Some welcoming text</h1>
          <h3>Some description</h3>
        </div>
      </>
  );
}
