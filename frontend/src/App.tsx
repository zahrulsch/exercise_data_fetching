import { useEffect, useState } from "react"
import { GetProductsResponse, getProducts } from "./apis/http"
import Loader from "./components/Loader"
import CardProduct from "./components/CardProduct"
import Pagination from "./components/Pagination"

function App() {
    const [params, setParams] = useState({
        size: 10,
        page: 1,
    })
    const [isLoading, setIsLoading] = useState(false)
    const [products, setProducts] = useState<GetProductsResponse | null>(null)

    useEffect(() => {
        setIsLoading(true)
        getProducts(params)
            .then(setProducts)
            .finally(() => setIsLoading(false))
    }, [params])

    if (products) {
        return (
            <div
                css={{
                    maxWidth: "967px",
                    margin: "0 auto",
                    paddingBlock: 20,
                }}
            >
                <Pagination
                    current={params.page}
                    pageSize={10}
                    totalData={products.count}
                    onChange={(page) => setParams((par) => ({ ...par, page }))}
                />
                {!isLoading ? (
                    <div
                        css={{
                            display: "flex",
                            flexWrap: "wrap",
                            gap: "10px 10px",
                            marginBottom: 10,
                        }}
                    >
                        {products.data.map((product) => (
                            <CardProduct
                                key={product.id}
                                imageURL={product.imageUrl}
                                title={product.name}
                            />
                        ))}
                    </div>
                ) : (
                    <div
                        css={{
                            width: "100%",
                            display: "flex",
                            justifyContent: "center",
                            alignItems: "center",
                            minHeight: "500px",
                        }}
                    >
                        <Loader width="min-content" height="min-content" />
                    </div>
                )}
            </div>
        )
    }

    return (
        <div
            css={{
                display: "flex",
                width: "100%",
                height: "100vh",
                justifyContent: "center",
                alignItems: "center",
            }}
        >
            <Loader width="min-content" height="min-content" />
        </div>
    )
}

export default App
