import clsx from "clsx";
import Heading from "@theme/Heading";
import styles from "./styles.module.css";

const FeatureList = [
  {
    title: "⚡ Lightweight & Fast",
    description: (
      <>
        A small, fast Go-based CLI tool designed to make sending HTTP requests simple and 
        efficient without any extra dependencies or setup.
      </>
    ),
  },
  {
    title: "💻 Simple CLI Requests",
    description: (
      <>
        Quickly send HTTP requests directly from your terminal using intuitive commands that 
        let you test APIs in seconds.
      </>
    ),
  },
  {
    title: "🎨 Prettified Colorful Output",
    description: (
      <>
        Response data is displayed in a readable, color-coded format that highlights JSON 
        structure, headers, and status codes for easy inspection.
      </>
    ),
  },
];

function Feature({title, description}) {
  return (
    <div className={clsx("col col--4")}>
      <div className="text--center padding-horiz--md">
        <Heading as="h2">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomeFeatures() {
  return (
    <section className={styles.featuresSection}>
      <div className="container">
        <Heading as="h1" className={styles.featuresHeading}>Features</Heading>
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
