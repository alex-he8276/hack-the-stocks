import * as React from "react";
import { Player } from '@lottiefiles/react-lottie-player';

const Loading = () => {
  return (
    <div className="flex w-full h-full justify-center items-center my-10 2xl:mt-20">
        <Player
            autoplay
            loop
            src="https://assets8.lottiefiles.com/packages/lf20_raiw2hpe.json"
            style={{ height: '300px', width: '300px' }}
        ></Player>
    </div>
  );
};

export default Loading;
