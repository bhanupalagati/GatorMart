import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import {Product} from '../interfaces/product.interface';

@Component({
  selector: 'app-card',
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.scss']
})
export class CardComponent implements OnInit {
  @Input() product: Product;
  constructor(private router: Router) { }

  ngOnInit(): void {
    // Get product as in put
  }

  detailedView(id) {
    this.router.navigateByUrl('/product/'+id);
    // Route to the details page
  }

}
