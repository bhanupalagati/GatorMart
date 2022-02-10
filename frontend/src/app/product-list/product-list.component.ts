import { Component, OnInit } from '@angular/core';
import { Product, ProductResponse } from '../interfaces/product.interface';
import { productResp } from '../mocks/products.mock';
import { HttpClient } from '@angular/common/http';
import { ProductsService } from '../services/products.service';

@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.scss']
})

export class ProductListComponent implements OnInit {
  productResp: Product;
  constructor(private productsService: ProductsService) { }

  ngOnInit(): void {
    this.productsService.getProductList().subscribe(res => {
      this.productResp = res;
      console.log(this.productResp);
      
    });
  }

}