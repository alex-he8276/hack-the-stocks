import * as React from 'react'
import Layout from "./layout/layout";
import Search from "./util/search";
import Button from "./util/button";
import SearchIcon from "@mui/icons-material/Search";
import Graph from "./components/graph/graph";

const Page = () => {
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
          { loading? <>pog</> : Conditional()}
        </div>
      </div>
    </Layout>
    )
}