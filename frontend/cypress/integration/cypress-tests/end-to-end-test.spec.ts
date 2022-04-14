// Objective: Complete end to end test Cypress video
/// <reference types ="cypress" />

describe('Login', function () {
    it('Sign up functionality', function() {
        cy.visit('http://localhost:4200/signup')
        })
    it('Login function', function() {
        //Email
        cy.get('input').eq(0).type('user@gatormart.com') 
        //Password
        cy.get('input').eq(1).type('Password1$') 
    })
    it('Submit Button', function() {
        cy.get('button').click()       
    })
})