import "./App.css";
import Layout from "./components/layout/layout";
import Search from "./components/util/search";
import Button from "./components/util/button";
import SearchIcon from "@mui/icons-material/Search";
import Graph from "./components/graph/graph";
function App() {
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
            <Search />
            <Button>
              Search <SearchIcon />
            </Button>
          </div>
          <Graph  className="h-28"/>
        </div>
      </div>
    </Layout>
  );
}

export default App;
