import Heading from "@theme/Heading";
import styles from "./styles.module.css";

export default function HomeWhatIsSAReq() {
  return (
    <section className={styles.sareqSection}>
      <div className="container">
        <Heading as="h1" className={styles.sareqHeading}>What is SAReq?</Heading>
        <p>
          SAReq, short for "Send A Request", is a command-line HTTP client to help 
          developers and testers send HTTP requests and analyze responses directly from the 
          terminal. It aims to provide a lightweight yet powerful and intuitive environment 
          that minimizes overhead and complexity of the application while maintaining 
          essential developer functionality.
        </p>
        <p>SAReq is built for:</p>
        <ul className={styles.sareqBulletList}>
          <li>
            Developers who are just getting started and want to experiment with HTTP 
            requests in a straightforward way.
          </li>
          <li>
            Developers who prefer command-line tools for quick testing and debugging of APIs.
          </li>
        </ul>
      </div>
    </section>
  );
}
