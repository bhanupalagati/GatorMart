import { Component, OnInit } from '@angular/core';
import { ProductResponse } from '../interfaces/product.interface';
import { productResp } from '../mocks/products.mock';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.scss']
})

export class ProductListComponent implements OnInit {
  productResp: ProductResponse;
  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    this.productResp = productResp;
    this.getConfig().subscribe(res => console.log(res));
  }

  getConfig() {
    return this.http.get("http://127.0.0.1:8000/products");
  }
}