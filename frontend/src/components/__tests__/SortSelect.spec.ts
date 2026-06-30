import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import SortSelect from '../SortSelect.vue'

describe('SortSelect', () => {
  it('emits the selected value to the parent', async () => {
    const wrapper = mount(SortSelect)

    await wrapper.get('select').setValue('asc')

    expect(wrapper.emitted('update:modelValue')).toBeTruthy()
    expect(wrapper.emitted('update:modelValue')?.[0]).toEqual(['asc'])
  })
})
