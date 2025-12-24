import Heading from "@theme/Heading";
import Link from "@docusaurus/Link";
import styles from "./styles.module.css";

export default function HomeContributionsDevelopment() {
  return (
    <section className={styles.contributionsDevelopmentSection}>
      <div className="container">
        <Heading as="h1">Contributions & Future Development</Heading>
        <p>
          SAReq is open-source and hosted on GitHub. Contributions are welcome in the form 
          of bug reports, bug fixes and code improvements, feature requests and documentation 
          updates. Future development plans include adding more cool features, usability 
          enhancements and additional CLI options.
        </p>
        <p>
          For more details on how to contribute and future roadmap, please refer to the{" "}
          <Link to="/docs/contributing">Contributing</Link> guide or the{" "}
          <Link to="/docs/roadmap">Roadmap</Link> document.
        </p>
      </div>
    </section>
  );
}
