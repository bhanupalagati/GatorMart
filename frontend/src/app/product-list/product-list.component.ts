import { Component, OnInit } from '@angular/core';
import { ProductResponse } from '../interfaces/product.interface';
import { productResp } from '../mocks/products.mock';

@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.scss']
})
export class ProductListComponent implements OnInit {
  productResp: ProductResponse;
  constructor() { }

  ngOnInit(): void {
    this.productResp = productResp;
  }

}
