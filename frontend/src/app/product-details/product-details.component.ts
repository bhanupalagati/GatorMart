import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import { EditFormService } from '../ad-form/edit-form.service';
import { DeletePopUpComponent } from '../delete-pop-up/delete-pop-up.component';
import { Product } from '../interfaces/product.interface';
import { product } from '../mocks/products.mock';
import { ProductsService } from '../services/products.service';
import { GoogleMapsModule } from '@angular/google-maps';

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
  constructor(private activatedRoute: ActivatedRoute,
    private productService: ProductsService, public dialog: MatDialog,
    private editFormService: EditFormService, private router: Router) { }

  ngOnInit(): void {
    this.id = this.activatedRoute.snapshot.paramMap.get("id");
    this.fetchProduct();
  }

  editForm() {
    this.editFormService.setFormData({...this.product});
    console.log(this.product);
    
    this.router.navigate(['/edit']);
  }

  fetchProduct() {
    this.productService.getProductDetails(this.id).subscribe(res => {
      this.product = res;
      this.images = this.product.images.split(",");
      this.loading = false;
    });
  }

  openDelete() {
    this.dialog.open(DeletePopUpComponent, { data: this.product });
  }

  //MapCode - Start
  mapOptions: google.maps.MapOptions = {
    center: { lat: Number(product.location_lat), lng: Number(product.location_long) },
    zoom: 14
  };
  marker = { position: { lat: Number(product.location_lat), lng: Number(product.location_long) } };
  //MapCode - End
}
