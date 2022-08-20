import * as React from 'react'

const Button = ({onClick, children})  => {
    return(
        <button className="bg-green-300 px-5 rounded-md hover:bg-green-200" onClick={() => onClick()}>{children}</button>
    )
}

export default Button