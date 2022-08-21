import * as React from "react";
import { ResponsiveLine } from "@nivo/line";

const Graph = ({ data , sentiment , stockName}) => {
  return (
    <>
      <div className="relative outline-green-300 outline rounded-2xl my-5 2xl:mt-18 px-5 w-[48rem] h-[20.5rem] 2xl:h-[38rem] 2xl:w-[62rem]">
        <div className="flex w-full justify-center mt-1"><h1>{stockName}</h1></div>
        <div className="absolute w-[48rem] h-[19rem] 2xl:h-[35rem] 2xl:w-[62rem] z-[4]">
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
            axisBottom={{
              orient: "bottom",
              tickSize: 5,
              tickPadding: 5,
              tickRotation: 0,
              legend: "date",
              legendOffset: 36,
              legendPosition: "middle",
            }}
            axisRight={null}
            axisLeft={{
              orient: "left",
              tickSize: 5,
              tickPadding: 5,
              tickRotation: 0,
              legend: "price",
              legendOffset: -50,
              legendPosition: "middle",
            }}
            colors={{ scheme: "set1" }}
            pointSize={10}
            pointColor={{ theme: "background" }}
            pointBorderWidth={2}
            pointBorderColor={{ from: "serieColor" }}
            pointLabelYOffset={-12}
            useMesh={true}
            legends={[
              {
                anchor: "top-left",
                direction: "column",
                justify: false,
                translateX: 0,
                translateY: -35,
                itemsSpacing: 0,
                itemDirection: "left-to-right",
                itemWidth: 80,
                itemHeight: 20,
                itemOpacity: 0.75,
                symbolSize: 12,
                symbolShape: "circle",
                symbolBorderColor: "rgba(0, 0, 0, .5)",
              },
            ]}
          />
        </div>
        <div className="absolute w-[48rem] h-[19rem] 2xl:h-[35rem] 2xl:w-[62rem] z-[5]">
          <ResponsiveLine
            data={sentiment}
            margin={{ top: 50, right: 110, bottom: 50, left: 60 }}
            xScale={{ type: "point" }}
            yScale={{
              type: "linear",
              min: 0,
              max: 100,
              stacked: true,
              reverse: false,
            }}
            enableGridY={false}
            yFormat=" >-.2f"
            axisTop={null}
            axisLeft={null}
            axisBottom={null}
            axisRight={{
              orient: "right",
              tickSize: 5,
              tickPadding: 5,
              tickRotation: 0,
              legend: "sentiment",
              legendOffset: 50,
              legendPosition: "middle",
            }}
            colors={{ scheme: "set2" }}
            pointSize={10}
            pointColor={{ theme: "background" }}
            pointBorderWidth={2}
            pointBorderColor={{ from: "serieColor" }}
            pointLabelYOffset={-12}
            useMesh={false}
            legends={[
              {
                anchor: "top-left",
                direction: "column",
                justify: false,
                translateX: 90,
                translateY: -35,
                itemsSpacing: 0,
                itemDirection: "left-to-right",
                itemWidth: 80,
                itemHeight: 20,
                itemOpacity: 0.75,
                symbolSize: 12,
                symbolShape: "circle",
                symbolBorderColor: "rgba(0, 0, 0, .5)",
              },
            ]}
          />
        </div>
      </div>
    </>
  );
};

export default Graph;
