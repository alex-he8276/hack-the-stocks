import "./App.css";
import Layout from "./components/layout/layout";
import Search from "./components/util/search";
import Button from "./components/util/button";
import SearchIcon from "@mui/icons-material/Search";
import Graph from "./components/graph/graph";
import Loading from "./components/util/loading";
import { useState } from "react";
import * as happystockapi from "./happystockapi_pb";

function App() {
  const [loading, setLoading] = useState(false);
  const [search, setSearch] = useState(false);
  const [stockPriceData, setStockPriceData] = useState([]);
  const [sentimentData, setSentimentData] = useState([]);
  const [ticker, setTicker] = useState("");
  const [companyName, setCompanyName] = useState("");
  const [graphLabel, setGraphLabel] = useState("");

  const hitStockApi = (ticker) => {
    const apiUrl = `http://127.0.0.1:5000/stock_price?ticker=${ticker}`;
    fetch(apiUrl)
      .then((response) => response.arrayBuffer())
      .then((data) => {
        let myData = happystockapi.listStockPrice.deserializeBinary(data);
        let dataObj = myData.toObject();
        myData = dataObj.pricelistList.map((item) => {
          return {
            x: new Date(item.date.seconds * 1000).toLocaleDateString("en-US"),
            y: item.price,
          };
        });
        const result = [];
        result.push({
          id: "Stock Price",
          data: myData,
        });
        setStockPriceData(result);
      })
      .catch((err) => console.log(err));
  };

  const hitSentimentApi = (ticker) => {
    /* placeholder */
    const apiUrl = `http://localhost:8080/api/stock/${ticker}`;
    fetch(apiUrl)
      .then((response) => response.arrayBuffer())
      .then((data) => {
        let myData = happystockapi.listStockSentiment.deserializeBinary(data);
        let dataObj = myData.toObject();
        console.log(dataObj);
        myData = dataObj.sentimentlistList.map((item) => {
          return {
            x: new Date(item.date.seconds * 1000).toLocaleDateString("en-US"),
            y: item.sentiment,
          };
        });
        console.log(myData);
        const result = [];
        result.push({
          id: "Sentiment",
          data: myData,
        });
        setSentimentData(result);
        setLoading(false);
      })
      .catch((err) => console.log(err));
  };

  const buttonClick = () => {
    if (companyName !== graphLabel && ticker) {
      setLoading(true);
      setSearch(true);
      setStockPriceData([]);
      setGraphLabel(companyName);
      console.log(ticker);
      hitStockApi(ticker);
      hitSentimentApi(ticker);
    }
  };

  return (
    <>
      <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <link rel="preconnect" href="https://fonts.googleapis.com"></link>
        <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin></link>
        <link href="https://fonts.googleapis.com/css2?family=Comfortaa:wght@500&display=swap" rel="stylesheet"></link>
      </head>
      <Layout>
        <div className="w-2/3 bg-slate-50 mx-auto my-8 rounded-3xl shadow-2xl">
          <div className="flex flex-col items-center">
            <div className="justify-center flex">
              <h1 className="text-5xl 2xl:text-6xl mt-10 2xl:ml-15 font-comfortaa">
                Happy Stocks
              </h1>
            </div>
            <div className="mt-5 2xl:mt-7 flex space-x-2">
              <Search setValue={setTicker} setInputValue={setCompanyName} />
              <Button onClick={buttonClick}>
                Search <SearchIcon />
              </Button>
            </div>
            {loading ? (
              <Loading />
            ) : search ? (
              <>
                <Graph
                  data={stockPriceData}
                  sentiment={sentimentData}
                  stockName={graphLabel}
                />
              </>
            ) : (
              <></>
            )}
          </div>
        </div>
      </Layout>
    </>
  );
}

export default App;
