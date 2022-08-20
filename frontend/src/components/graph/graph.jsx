import * as React from "react";
import { ResponsiveLine } from "@nivo/line";

const fakeData = {
  listStockPrice: [
    { date: 1660276800, stockPrice: 291.29, tickerName: "MSFT" },
    { date: 1660536000, stockPrice: 292.84, tickerName: "MSFT" },
    { date: 1660622400, stockPrice: 292.08, tickerName: "MSFT" },
    { date: 1660708800, stockPrice: 291.32, tickerName: "MSFT" },
    { date: 1660795200, stockPrice: 290.17, tickerName: "MSFT" },
  ],
};

const Graph = () => {
  const myData = fakeData.listStockPrice.map((item) => ({
    x: new Date(item.date * 1000).toLocaleDateString("en-US"),
    y: item.stockPrice,
  }));
  const data = [];
  data.push({
    id: fakeData.listStockPrice[0].tickerName,
    data: myData,
  });
  console.log(data);
  return (
    <>
    <div className="outline-green-300 outline rounded-2xl mt-8 2xl:mt-18 px-5 w-[48rem] h-[19rem] 2xl:h-[35rem] 2xl:w-[62rem]">
      <ResponsiveLine
        data={data}
        margin={{ top: 50, right: 110, bottom: 50, left: 60 }}
        xScale={{ type: "point" }}
        yScale={{
          type: "linear",
          min: "auto",
          max: "auto",
          stacked: true,
          reverse: false,
        }}
        yFormat=" >-.2f"
        axisTop={null}
        axisRight={null}
        axisBottom={{
          orient: "bottom",
          tickSize: 5,
          tickPadding: 5,
          tickRotation: 0,
          legend: "transportation",
          legendOffset: 36,
          legendPosition: "middle",
        }}
        axisLeft={{
          orient: "left",
          tickSize: 5,
          tickPadding: 5,
          tickRotation: 0,
          legend: "price",
          legendOffset: -50,
          legendPosition: "middle",
        }}
        colors={{ scheme: 'set1' }}
        pointSize={10}
        pointColor={{ theme: "background" }}
        pointBorderWidth={2}
        pointBorderColor={{ from: "serieColor" }}
        pointLabelYOffset={-12}
        useMesh={true}
        legends={[
          {
            anchor: "bottom-right",
            direction: "column",
            justify: false,
            translateX: 100,
            translateY: 0,
            itemsSpacing: 0,
            itemDirection: "left-to-right",
            itemWidth: 80,
            itemHeight: 20,
            itemOpacity: 0.75,
            symbolSize: 12,
            symbolShape: "circle",
            symbolBorderColor: "rgba(0, 0, 0, .5)",
            effects: [
              {
                on: "hover",
                style: {
                  itemBackground: "rgba(0, 0, 0, .03)",
                  itemOpacity: 1,
                },
              },
            ],
          },
        ]}
      />
    </div>
    </>
  );
};

export default Graph;
