import { Component, Input, OnInit } from '@angular/core';
import {Product} from '../interfaces/product.interface';

@Component({
  selector: 'app-card',
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.scss']
})
export class CardComponent implements OnInit {
  @Input() product: Product;
  constructor() { }

  ngOnInit(): void {
    // Get product as in put
  }

  detailedView(event) {
    // Route to the details page
  }

}
