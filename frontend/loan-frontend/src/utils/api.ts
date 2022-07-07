export const fetcher = (path: string) => {
  fetch(`http://0.0.0.0:1317${path}`).then((res) => res.json()).then((data) => data)
};
