import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateAddComponent } from './create-add.component';

describe('CreateAddComponent', () => {
  let component: CreateAddComponent;
  let fixture: ComponentFixture<CreateAddComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateAddComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
