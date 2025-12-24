import Tabs from "@theme/Tabs";
import TabItem from "@theme/TabItem";
import CodeBlock from "@theme/CodeBlock";
import styles from "./styles.module.css";

export default function InstallChecksumVerification() {
  return (
    <div className={styles.installChecksumVerificationCard}>
      <Tabs className={styles.installChecksumVerificationTabs}>
        <TabItem value="Windows" label="Windows" default>
          <CodeBlock language="powershell">
            Get-FileHash -Algorithm SHA256 NAME_OF_DOWNLOADED_ARCHIVE.zip
          </CodeBlock>
        </TabItem>
        <TabItem value="Linux" label="Linux">
          <CodeBlock language="bash">
            sha256sum NAME_OF_DOWNLOADED_ARCHIVE.tar.gz
          </CodeBlock>
        </TabItem>
        <TabItem value="macOS" label="macOS">
          <CodeBlock language="bash">
            shasum -a 256 NAME_OF_DOWNLOADED_ARCHIVE.tar.gz
          </CodeBlock>
        </TabItem>
      </Tabs>
    </div>
  );
}
