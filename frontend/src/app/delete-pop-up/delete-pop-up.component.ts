import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { Product } from '../interfaces/product.interface';
import { ProductsService } from '../services/products.service';

@Component({
  selector: 'app-delete-pop-up',
  templateUrl: './delete-pop-up.component.html',
  styleUrls: ['./delete-pop-up.component.scss']
})
export class DeletePopUpComponent implements OnInit {

  constructor(public dialogRef: MatDialogRef<DeletePopUpComponent>,
              @Inject(MAT_DIALOG_DATA) public data: Product,
              private productsService: ProductsService,
              private router: Router) { }

  ngOnInit(): void {
  }

  deleteAd(id) {
    this.productsService.deleteProduct(this.data.ID).subscribe(res => {
      this.router.navigate(['/']);
    });
    this.dialogRef.close();
  }

}
