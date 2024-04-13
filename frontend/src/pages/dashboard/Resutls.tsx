import { useEffect } from "react";

const Results = ({ pageTitle }: { pageTitle: string }) => {
  useEffect(() => {
    document.title = pageTitle;
  }, [pageTitle]);
  return <div>Results</div>;
};

export default Results;
