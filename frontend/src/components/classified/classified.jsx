import * as React from "react";
import { TwitterTweetEmbed } from "react-twitter-embed";

const TweetCard = ({ classification, id }) => {
  const classificationColor =
    classification === "Positive" ? "text-green-500" : "text-red-500";
  const borderColor =
    classification === "Positive" ? "outline-green-300" : "outline-red-300";
  return (
    <div className="my-5 w-full">
      <div
        className={`w-full flex flex-col ${borderColor} outline rounded-2xl p-3 items-center`}
      >
        <div className="flex justify-center">
          <h4 className="font-comfortaa font-bold">
            Classification:{" "}
            <span className={`${classificationColor}`}>{classification}</span>
          </h4>
        </div>
        <TwitterTweetEmbed tweetId={id} />
        
       
      </div>
    </div>
  );
};

const ClassifiedTweets = ({tweetList}) => {
  return (
    <div className="w-2/3 bg-slate-50 mx-auto mb-8 rounded-3xl shadow-2xl">
      <div className="flex justify-center">
        <div className="flex flex-col items-center">
          <h1 className="text-5xl h-full 2xl:text-6xl mt-10 2xl:ml-15 font-comfortaa">
            Classification Examples
          </h1>
          {tweetList.map((item) => (
            <TweetCard classification={item.classification} id={item.id} />
          ))}
        </div>
      </div>
    </div>
  );
};

export default ClassifiedTweets;
