import { Component, OnInit } from '@angular/core';
import { Product } from '../interfaces/product.interface';
import { ProductsService } from '../services/products.service';

@Component({
  selector: 'app-my-ads',
  templateUrl: './my-ads.component.html',
  styleUrls: ['./my-ads.component.scss']
})
export class MyAdsComponent implements OnInit {
  productResp: Product[];
  loading = true;
  constructor(private productsService: ProductsService) { }

  ngOnInit(): void {
    this.productsService.getProductsByUser().subscribe(res => {
      this.productResp = res;
      this.loading = false;
    });
  }

}
