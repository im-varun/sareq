import Heading from "@theme/Heading";
import Tabs from "@theme/Tabs";
import TabItem from "@theme/TabItem";
import CodeBlock from "@theme/CodeBlock";
import Link from "@docusaurus/Link";
import styles from "./styles.module.css";

export default function HomeQuickInstall() {
  return (
    <section className={styles.quickInstallSection}>
      <div className="container">
        <Heading as="h1">Quick Install</Heading>
        <p className={styles.quickInstallDescription}>
          SAReq can be installed via the command line or by downloading the official binary 
          from GitHub Releases.
        </p>
        <div className={styles.quickInstallCard}>
          <Tabs className={styles.quickInstallTabs}>
            <TabItem value="go install" label="go install" default>
              <p>This assumes you have Go installed on your system.</p>
              <CodeBlock language="powershell">
                go install github.com/im-varun/sareq@latest
              </CodeBlock>
              <p className={styles.quickInstallFooterText}>
                Make sure your Go bin directory (usually <code>$GOPATH/bin</code> or 
                <code>$HOME/go/bin</code>) is in the system PATH. Once you have it, you are
                good to go! Head over to the{" "}
                <Link to="/docs/intro/quickstart">Quick Start</Link> guide to get started with SAReq.
              </p>
            </TabItem>
            <TabItem value="Binary" label="Binary">
              <p>
                Get the latest release for SAReq binary from{" "}
                <a href="https://github.com/im-varun/sareq/releases" 
                  target="_blank" 
                  rel="noopener noreferrer">
                  GitHub Releases
                </a>
              </p>
              <p className={styles.quickInstallFooterText}>
                Once downloaded, unpack the binary, add it to your PATH and you are good to 
                go! Head over to the{" "}
                <Link to="/docs/intro/quickstart">Quick Start</Link> guide to get started with SAReq.
              </p>
            </TabItem>
            <TabItem value="More" label="More">
              <p>
                More installation options will be added in the future. Stay tuned!
              </p>
            </TabItem>
          </Tabs>
        </div>
      </div>
    </section>
  );
}
