import isValidProp from "@emotion/is-prop-valid"
import styled from "@emotion/styled"

export type CardProductProps = {
    children?: React.ReactNode
    imageURL?: string
    title?: string
}

const _Card = styled("div", {
    shouldForwardProp: (propName) => isValidProp(propName),
})<CardProductProps>(() => ({
    width: 200,
    display: "flex",
    padding: 10,
    flexDirection: "column",
    justifyContent: "center",
    rowGap: 10,
    ":hover": {
        cursor: "pointer",
        "& h5": {
            color: "#1d4ed8",
        },
    },
    "& img": {
        width: "100%",
        height: "auto",
    },
    "& h5": {
        fontWeight: 500,
        transition: "300ms ease",
    },
}))

function CardProduct(props: CardProductProps): JSX.Element {
    return (
        <_Card {...props}>
            <img title={props.title} alt={props.title} src={props.imageURL} />
            <h5>{props.title}</h5>
        </_Card>
    )
}

export default CardProduct
