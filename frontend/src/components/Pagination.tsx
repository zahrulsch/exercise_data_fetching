import { css } from "@emotion/react"
import { useMemo } from "react"

export type PaginationProps = {
    children?: React.ReactNode
    pageSize?: number
    totalData?: number
    current?: number
    onChange?: (page: number) => void
}

const buttonCss = css`
    padding: 7px 12px;
    display: inline-flex;
    border: none;
    outline: none;
    background-color: #2563eb;
    color: #fff;
    border-radius: 5px;
    cursor: pointer;
    transition: 300ms ease;
    border: 1px solid transparent;
    &:hover {
        transition: 100ms ease;
        color: #2563eb;
        background-color: #fff;
        border: 1px solid #2563eb;
    }
    &.active {
        background-color: #eff6ff;
        color: #1d4ed8;
    }
`

function Pagination(props: PaginationProps): JSX.Element {
    const pageList = useMemo(() => {
        if (!props.pageSize && !props.totalData) return []
        const total = Math.ceil(props.totalData! / props.pageSize!)
        return Array.from(Array(total).keys())
    }, [props.pageSize, props.totalData])

    return (
        <div
            css={{
                display: "flex",
                columnGap: 10,
            }}
        >
            {pageList
                .map((n) => n + 1)
                .map((page) => (
                    <button
                        className={props.current == page ? "active" : ""}
                        key={page}
                        css={buttonCss}
                        onClick={() => props.onChange?.(page)}
                    >
                        {page}
                    </button>
                ))}
        </div>
    )
}

export default Pagination
