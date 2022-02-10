import { Component, Input, OnInit } from '@angular/core';
import { Product } from '../interfaces/product.interface';
import { product } from '../mocks/products.mock';

@Component({
  selector: 'app-product-details',
  templateUrl: './product-details.component.html',
  styleUrls: ['./product-details.component.scss']
})
export class ProductDetailsComponent implements OnInit {
  product: Product
  images: string[]
  constructor() { }

  ngOnInit(): void {
    this.product = product;
    this.images = product.images.split(",");
  }

  // Construct an API call to get product details
  

}
