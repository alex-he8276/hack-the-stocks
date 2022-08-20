import "./App.css";
import Layout from "./components/layout/layout";
import Search from "./components/util/search";
import Button from "./components/util/button";
import SearchIcon from "@mui/icons-material/Search";
import Graph from "./components/graph/graph";
import { useState } from "react";
import * as happystockapi from "./happystockapi_pb";

function App() {
  const [loading, setLoading] = useState(false);
  const [search, setSearch] = useState(false);
  const [stockPriceData, setStockPriceData] = useState([]);
  const [ticker, setTicker] = useState("")
  const [companyName, setCompanyName] = useState("")

  const hitApi = (ticker) => {
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
          id: ticker,
          data: myData,
        });
        setStockPriceData(result);
        setLoading(false);
      });
  };

  const buttonClick = () => {
    setLoading(true);
    setSearch(true);
    setStockPriceData([]);
    console.log(ticker)
    hitApi(ticker);
  };

  return (
    <Layout>
      <div className="w-2/3 bg-slate-50 mx-auto my-8 rounded-3xl shadow-2xl">
        <div className="flex flex-col items-center">
          <div className="justify-center flex">
            <h1 className="text-5xl 2xl:text-6xl mt-10 2xl:ml-15">
              Happy Stocks
            </h1>
          </div>
          <div className="mt-5 2xl:mt-7 flex space-x-2">
            <Search value={ticker} inputValue={companyName} setValue={setTicker} setInputValue={setCompanyName}/>
            <Button onClick={buttonClick}>
              Search <SearchIcon />
            </Button>
          </div>
          {loading ? (
            <>pog</>
          ) : search ? (
            <Graph className="h-28" data={stockPriceData} />
          ) : (
            <></>
          )}
        </div>
      </div>
    </Layout>
  );
}

export default App;
