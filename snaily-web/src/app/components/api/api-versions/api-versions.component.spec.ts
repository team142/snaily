import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ApiVersionsComponent } from './api-versions.component';

describe('ApiVersionsComponent', () => {
  let component: ApiVersionsComponent;
  let fixture: ComponentFixture<ApiVersionsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ApiVersionsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ApiVersionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
