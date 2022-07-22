import '@angular/localize/init';
import { moduleMetadata } from '@storybook/angular';
import { Meta, Story } from '@storybook/angular/types-6-0';
import { KtbDateInputModule } from '../../client/app/_components/ktb-date-input/ktb-date-input.module';
import { KtbDatetimePickerComponent } from '../../client/app/_components/ktb-date-input/ktb-datetime-picker.component';

export default {
  title: 'Components/Date Input',
  component: KtbDatetimePickerComponent,
  decorators: [
    moduleMetadata({
      imports: [KtbDateInputModule],
    }),
  ],
} as Meta;

const template: Story<KtbDatetimePickerComponent> = (args: KtbDatetimePickerComponent) => ({
  props: args,
});

export const date = template.bind({});
date.args = {
  timeEnabled: false,
};
