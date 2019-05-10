import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ApiIssuesComponent } from './api-issues.component';

describe('ApiIssuesComponent', () => {
  let component: ApiIssuesComponent;
  let fixture: ComponentFixture<ApiIssuesComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ApiIssuesComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ApiIssuesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
