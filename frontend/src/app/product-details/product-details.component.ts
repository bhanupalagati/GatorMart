import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute } from '@angular/router';
import { DeletePopUpComponent } from '../delete-pop-up/delete-pop-up.component';
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
  constructor(private activatedRoute: ActivatedRoute, private productService: ProductsService, public dialog: MatDialog) { }

  ngOnInit(): void {
    this.id = this.activatedRoute.snapshot.paramMap.get("id");
    this.fetchProduct();
  }

  fetchProduct() {
    this.productService.getProductDetails(this.id).subscribe(res => {
      this.product = res;
      this.images = this.product.images.split(",");
      this.loading = false;
    });
  }
  
  openDelete() {
    this.dialog.open(DeletePopUpComponent, {data: this.product});
  }
}
