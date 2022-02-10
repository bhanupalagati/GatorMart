import { Product, ProductResponse } from '../interfaces/product.interface';

export let product: Product = {
    title: 'Iphone 12 Pro Max',
    secondary_title: 'Black 4GB 256GB',
    imageUrl: "https://www.gizmochina.com/wp-content/uploads/2021/09/iPhone-13-all-colors-featured-b-1024x606.jpeg",
    price: 350,
    simple_desc: "This is a great device with 4Gb ram and 256Gb rom. Rarely used and everything is working smooth",
    description: "This is a complete description of this product and I am trying my best to use fillers here however, based on the scenario that is not the case and I am still going, going, going, and going forever.",
    city: "Gainesville",
    state: "FL",
    location_lat: "28.514179",
    location_long: "-81.367172",
    posted_by: "UID1",
    posted_date: "1/13/2022, 2:45:32 PM",
    condition: "Used",
    age: 4, 
    status: "Active",
    target: "student",
    category: "Electronics",
    images: "https://images.unsplash.com/photo-1601784551446-20c9e07cdbdb?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1867&q=80,https://images.unsplash.com/photo-1601784551446-20c9e07cdbdb?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1867&q=80,https://www.gizmochina.com/wp-content/uploads/2021/09/iPhone-13-all-colors-featured-b-1024x606.jpeg"
}

export let productResp: ProductResponse = {
    count: 5,
    next_page: -1,
    previous_page: -1,
    products: [{...product}, {...product}, {...product}, {...product}, {...product}, {...product}, {...product}, {...product}]
}