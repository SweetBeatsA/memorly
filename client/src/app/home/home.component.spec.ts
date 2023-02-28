import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatCardModule } from '@angular/material/card';

import { HomeComponent } from './home.component';

describe('HomeComponent', () => {
    let component: HomeComponent;
    let fixture: ComponentFixture<HomeComponent>;
  
    beforeEach(async () => {
      await TestBed.configureTestingModule({
        declarations: [ HomeComponent ],
        imports: [ MatToolbarModule, MatCardModule ]
      })
      .compileComponents();
    });
  
    beforeEach(() => {
      fixture = TestBed.createComponent(HomeComponent);
      component = fixture.componentInstance;
      fixture.detectChanges();
    });
  
    it('should display "Welcome to our CEN3031 Project!"', () => {
        const fixture = TestBed.createComponent(HomeComponent);
        const app = fixture.nativeElement.querySelector('div');
        expect(app.textContent).toContain('Welcome to our CEN3031 Project!');
      });
  });

