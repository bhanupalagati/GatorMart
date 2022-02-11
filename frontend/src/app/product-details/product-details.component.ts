import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Product } from '../interfaces/product.interface';
import { product } from '../mocks/products.mock';
import { ProductsService } from '../services/products.service';

@Component({
  selector: 'app-product-details',
  templateUrl: './product-details.component.html',
  styleUrls: ['./product-details.component.scss']
})
export class ProductDetailsComponent implements OnInit {
  product: Product;
  images: string[];
  id: string;
  loading = true;
  constructor(private activatedRoute: ActivatedRoute, private productService: ProductsService) { }

  ngOnInit(): void {
    this.id = this.activatedRoute.snapshot.paramMap.get("id");
    this.fetchProduct();
  }

  fetchProduct() {
    this.productService.getProductDetails(this.id).subscribe(res => {
      this.product = res;
      this.images = product.images.split(",");
      this.loading = false;
    });
  }
  // Construct an API call to get product details
  

}
