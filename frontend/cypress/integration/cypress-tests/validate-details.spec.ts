/// <reference types="cypress" />

describe('Access', function () {
    beforeEach(() => {
        cy.visit('http://localhost:4200/')
    })

    it('Verify the details submitted by the user', function() {
        cy.get('mat-card').eq(2).click()
        cy.wait(500) 
        cy.get('div [class="carousel-caption"]').find('h3').contains('Apple iPad M1')
        cy.wait(500)
        cy.get('div [class="carousel-caption"]').find('p').contains('iPad Pro 12.9')
        cy.wait(500)
        cy.get('div [class="details_container__main_info"]').find('div').eq(3).contains(1200)
        cy.wait(500)
        cy.get('div [class="details_container__main_info"]').find('div').eq(4).contains('Miami, Florida')
        cy.wait(500)
    })
})