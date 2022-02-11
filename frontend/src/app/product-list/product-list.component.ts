import { Component, OnInit } from '@angular/core';
import { Product } from '../interfaces/product.interface';
import { ProductsService } from '../services/products.service';

@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.scss']
})

export class ProductListComponent implements OnInit {
  productResp: Product[];
  loading = true;
  constructor(private productsService: ProductsService) { }

  ngOnInit(): void {
    this.productsService.getProductList().subscribe(res => {
      this.productResp = res;
      this.loading = false;
    });
  }
}