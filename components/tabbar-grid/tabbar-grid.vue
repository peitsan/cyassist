<!--
      参数                说明                         类型                  默认值
      girdData       宫格里面的内容                    Array                  []
      columnNum         列数                          String                  4
      gutter          格子之间的间距，请带单位传入      String                  0
                      如：px、vw
      gutterBottom    格子上下的间距，请带单位传入      String                  0
                      如：px、vw
使用例子:
    <gird :girdData="girdData" :gutter="'1.563vw'" :gutterBottom="'1.563vw'" :columnNum="'4'">
      <template #content="acceptProps">
        <div class="card">
          <div class="icon">
            <i class="iconfont" :class="acceptProps.item.icon" />
          </div>
          <div class="chName">{{acceptProps.item.chName}}</div>
          <div class="enName">{{acceptProps.item.enName}}</div>
        </div>
      </template>
    </gird>

    说明：acceptProps是由slot传递过去的，其中传递了item、index这两个值,需要特别注意的是,如果需要自适应的话，
    插槽的内容不要设置width！(例如 <div class="card">不要设置宽度)
-->

<template>
  <div class="gird" :style="{paddingLeft:gutter}">
    <div
      class="gird-item"
      :style="{flexBasis:flexBasis, paddingRight:gutter,marginBottom:gutterBottom}"
      v-for="(item,index) in girdData"
      :key="index"
      id="girdChild"
    >
      <slot name="content" :item="item" :index="index"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: '',
  components: {},
  props: {
    girdData: {
      type: Array,
      default: () => {
        return []
      }
    },
    columnNum: {
      type: String,
      default: '4'
    },
    gutter: {
      type: String,
      default: '0'
    },
    gutterBottom: {
      type: String,
      default: '0'
    }
  },
  data () {
    return {}
  },
  computed: {
    flexBasis () {
      return 100 / Number(this.columnNum) + '%'
    }
  },
  mounted () {},
  methods: {}
}
</script>

<style lang="scss" scoped>
.gird {
  display: -webkit-box;
  display: -webkit-flex;
  display: flex;
  -webkit-flex-wrap: wrap;
  flex-wrap: wrap;
  width: 100%;
  box-sizing: border-box;
  .gird-item {
    box-sizing: border-box;
  }
}
</style>

