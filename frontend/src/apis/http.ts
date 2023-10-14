import { client } from "./client"

type GetProductsParams = {
    page?: number
    size?: number
}

export type GetProductsResponse = {
    count: number
    data: {
        catId: number
        countReview: number
        discountPercentage: number
        gaKey: string
        id: number
        imageUrl: string
        imageUrlLarge: string
        name: string
        original_price: string
        preorder: boolean
        price: string
        priceInt: number
        rating: number
        url: string
    }[]
}

async function getProducts(
    params: GetProductsParams = {
        page: 1,
        size: 10,
    }
) {
    const request = await client.get("/product", {
        params,
    })

    return request.data as GetProductsResponse
}

export { getProducts }
