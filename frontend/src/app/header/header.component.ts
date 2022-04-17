import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProductsService } from '../services/products.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  constructor(private productsService: ProductsService, private router: Router) { }
  userDetails: any;
  showFilter = false;
  ngOnInit(): void {
    this.userDetails = this.productsService.getUserData();
    console.log(this.userDetails);
    
  }

  navigateTo(url) {
    this.router.navigate(['/'+url]);
  }

  filterChange() {
    this.showFilter = !this.showFilter;
    this.productsService.filterMenuToggled.next(this.showFilter);
  }

  logout() {
    this.productsService.setCookies('token', '', 0);
    this.productsService.setCookies('userInfo', '', 0);
    this.router.navigate(['/']);
  }

}
