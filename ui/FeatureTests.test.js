import { Selector } from 'testcafe';

// eslint-disable-next-line no-undef
fixture('End to End Test').page('http://localhost:3000');

async function expectNewRoundPageSuccessful(t) {
  const newParticipantTextField = Selector('.new-round-participant-text-field');
  const addParticipantButton = Selector('#new-round-add-participant');
  const nextBuyerButton = Selector('#next-buyer-button');
  const currentCandidate = Selector('.current-candidate');

  return t.typeText(newParticipantTextField, 'John')
    .click(addParticipantButton)
    .typeText(newParticipantTextField, 'CamWyn')
    .click(addParticipantButton)
    .typeText(newParticipantTextField, 'Fergus')
    .click('#new-round-start-round')
    .expect(nextBuyerButton.innerText)
    .eql('NEXT BUYER')
    .expect(currentCandidate.innerText)
    .eql('Tom buys')
    .click(nextBuyerButton)
    .expect(currentCandidate.innerText)
    .eql('Geoff buys');
}

test('Complete shopping experience', async (t) => {
  await expectNewRoundPageSuccessful(t);
  await t.click('#new-round-button');
  await expectNewRoundPageSuccessful(t);
  await t.click('#other-rounds-button')
    .expect(Selector('.other-rounds-participant:nth-child(1)').innerText).eql('Bob')
    .click('.other-rounds-remove-round:nth-child(1)')
    .expect(Selector('.other-rounds-page-no-rounds').innerText)
    .eql('No other rounds');
  await t.click('#credits-link')
    .expect(Selector('#credits-title').innerText).eql('Icons and Images');
});
