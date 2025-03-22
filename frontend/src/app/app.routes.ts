import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { PetListComponent } from './pet-list/pet-list.component';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AppComponent } from './app.component';



export const routes: Routes = [
  { path: '', component: PetListComponent },
  { path: 'pets', component: PetListComponent }
];

@NgModule({
  declarations: [
  ],
  imports: [
    RouterModule.forRoot(routes),
    FormsModule,
    CommonModule
],
  providers: [],
  exports: [RouterModule, FormsModule, CommonModule]
  
})

export class AppRoutingModule { }
export class AppModule {}