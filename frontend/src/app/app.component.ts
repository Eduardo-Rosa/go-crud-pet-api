import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  template: `
    <div class="container">
      <router-outlet></router-outlet>
    </div>
  `,
  standalone: true,
  styleUrls: ['./app.component.css'],
  imports: [RouterOutlet],
  templateUrl: './app.component.html',
})
export class AppComponent {
  title = 'frontend';
}
