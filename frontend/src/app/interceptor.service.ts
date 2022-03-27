import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ProductsService } from './services/products.service';

@Injectable({
  providedIn: 'root'
})
export class InterceptorService implements HttpInterceptor {
  token: string;
  constructor(private productsService: ProductsService) { }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    this.token = this.productsService.getCookie();
    if (this.token) {
      const tokenizedReq = req.clone({ headers: req.headers.set('Authorization', this.token) });
      return next.handle(tokenizedReq);
    }
    return next.handle(req);
  }
}
