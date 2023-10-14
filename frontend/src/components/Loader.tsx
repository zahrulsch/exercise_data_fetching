import { css, keyframes } from "@emotion/react"

type NumberOrString = number | string
export type LoaderProps = {
    children?: React.ReactNode
    width?: NumberOrString
    height?: NumberOrString
    loaderText?: JSX.Element
    style?: React.CSSProperties
}

const loaderKeyframes = keyframes`
    from {
        transform: rotate(0deg);
        opacity: 1;
    }

    50% {
        opacity: 0.3;
    }

    to {
        opacity: 1;
        transform: rotate(360deg);
    }
`

const loaderCss = (props: LoaderProps) => {
    return css({
        display: "flex",
        flexDirection: "column",
        rowGap: 10,
        height: props.height || "100%",
        width: props.width || "100%",
        padding: 10,
        "& .circle-spinner": {
            height: 24,
            width: 24,
            borderRadius: 24,
            borderRight: "2px solid #60a5fa",
            borderBottom: "2px solid #60a5fa",
            borderLeft: "2px solid #dbeafe",
            borderTop: "2px solid #dbeafe",
            animation: `750ms ${loaderKeyframes} linear infinite`,
        },
    })
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const isString = (children: any): children is string =>
    typeof children == "string"

function Loader(props: LoaderProps): JSX.Element {
    return (
        <div css={loaderCss(props)} style={props.style}>
            <div className="circle-spinner"></div>
            {isString(props.children) ? (
                <span>{props.children}</span>
            ) : (
                props.children
            )}
        </div>
    )
}

export default Loader
