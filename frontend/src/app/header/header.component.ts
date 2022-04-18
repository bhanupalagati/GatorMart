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
    setTimeout(() => {
      if (!this.router.url.includes('signup')) {
        this.login();
      }
    }, 100);

    this.productsService.userData.subscribe(res => this.userDetails = res);
  }

  navigateTo(url) {
    this.router.navigate(['/' + url]);
  }

  filterChange() {
    this.showFilter = !this.showFilter;
    this.productsService.filterMenuToggled.next(this.showFilter);
  }

  login() {
    this.productsService.validateCookie().subscribe(res => {
      this.userDetails = res;
      this.productsService.userData.next(res);
      console.log(res);

    }, err => {
      console.table(err);

      this.router.navigate(['/']);
    }
    );
  }

  logout() {
    this.productsService.userData.next({})
    this.productsService.setCookies('token', '', 0);
    this.productsService.setCookies('userInfo', '', 0);
    this.router.navigate(['/']);
  }

}
