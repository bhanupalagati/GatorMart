export type Product = {
    title: string,
    secondary_title: string,
    imageUrl: string,
    price: number,
    simple_desc: string,
    description: string,
    city: string,
    state: string,
    location_lat: number,
    location_long: number,
    posted_by: string,
    posted_date: string,
    condition: "New" | "Like New" | "Used" | "Heavily Used" | "NA",
    age: number,
    status: "Deleted" | "Sold" | "Active",
    images: string[]
}

export type ProductResponse = {
    count: number,
    next_page: number,
    previous_page: number,
    products: Product[]
}