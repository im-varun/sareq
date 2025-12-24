import clsx from "clsx";
import Layout from "@theme/Layout";
import Heading from "@theme/Heading";
import HomeWhatIsSAReq from "@site/src/components/HomeWhatIsSAReq";
import HomeFeatures from "@site/src/components/HomeFeatures";
import HomeQuickInstall from "@site/src/components/HomeQuickInstall";
import HomeContributionsDevelopment from "@site/src/components/HomeContributionsDevelopment";
import styles from "./index.module.css";

function HomeHeader() {
  return (
    <header className={clsx("hero hero--primary", styles.heroBanner)}>
      <div className="container">
        <Heading as="h1" className="hero__title">
          SAReq /&gt;&gt;/
        </Heading>
        <p className="hero__subtitle">
          A modern, open-source HTTP client for the command line
        </p>
      </div>
    </header>
  );
}

export default function Home() {
  return (
    <Layout
      title="SAReq"
      description="A modern, open-source HTTP client for the command line">
      <HomeHeader />
      <main>
        <HomeWhatIsSAReq />
        <HomeFeatures />
        <HomeQuickInstall />
        <HomeContributionsDevelopment />
      </main>
    </Layout>
  );
}
