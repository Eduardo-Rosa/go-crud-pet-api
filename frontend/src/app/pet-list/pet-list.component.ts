import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { Routes } from '@angular/router';


interface Pet {
  id: number;
  name: string;
  species: string;
  breed: string;
  age: number;
  birthDate: string;
  ownerName: string;
}

@Component({
  selector: 'app-pet-list',
  standalone: true,
  imports: [HttpClientModule,
     CommonModule, 
     FormsModule, 
     RouterModule
    ],
  templateUrl: './pet-list.component.html',
  styleUrls: ['./pet-list.component.css']
})
export class PetListComponent implements OnInit {
  pets: Pet[] = [];
  pet: Pet = {
    id: 0,
    name: '',
    species: '',
    breed: '',
    age:0,
    birthDate: '',
    ownerName: ''
  };
  isEditing = false;

  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.loadPets();
  }

  loadPets(): void {
    this.http.get<Pet[]>('http://localhost:8080/pets').subscribe(
      (data) => this.pets = data,
      (error) => console.error('Erro ao carregar pets:', error)
    );
  }

  onSubmit(): void {
    if (this.isEditing) {
      this.http.put(`http://localhost:8080/pets/${this.pet.id}`, this.pet).subscribe(
        () => {
          this.loadPets();
          this.resetForm();
        },
        (error) => console.error('Erro ao atualizar pet:', error)
      );
    } else {
      this.http.post('http://localhost:8080/pets', this.pet).subscribe(
        () => {
          this.loadPets();
          this.resetForm();
        },
        (error) => console.error('Erro ao cadastrar pet:', error)
      );
    }
  }

  editPet(pet: Pet): void {
    this.pet = { ...pet };
    this.isEditing = true;
  }

  deletePet(id: number): void {
    this.http.delete(`http://localhost:8080/pets/${id}`).subscribe(
      () => this.loadPets(),
      (error) => console.error('Erro ao excluir pet:', error)
    );
  }

  resetForm(): void {
    this.pet = {
      id: 0,
      name: '',
      species: '',
      breed: '',
      age: 0,
      birthDate: '',
      ownerName: ''
    };
    this.isEditing = false;
  }
}
