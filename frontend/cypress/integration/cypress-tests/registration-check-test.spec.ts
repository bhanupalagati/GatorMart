//Objective: Test the registration functionality

/// <reference types="cypress" />

describe('Register an user normally', function () {
    it('Sign up functionality', function() {
        cy.visit('http://localhost:4200/signup')
        
    })

    it('Sign up form', function() {
        //Email
        cy.get('input')
        .eq(0).type('user@gatormart.com') 
        //Password
        cy.get('input')
        .eq(1).type('Password1$') 
        //Re-Type Password
        cy.get('input')
        .eq(2).type('Password1$') 
        //Profession
        cy.get('input')
        .eq(3).type('Student')
        //Name
        cy.get('input')
        .eq(4).type('Gatory')
        //Last name
        cy.get('input')
        .eq(5).type('Marty')
        //Last name
        cy.get('input')
        .eq(6).invoke('removeAttr','type').type('05-02-2011');
    })
    it('Submit registration form', function() {
        cy.get('button').click()       
    })
})


