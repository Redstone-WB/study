XML
- 안드로이드에서 View를 그리는데 사용되는 언어
- DSL (Domain Specific Language)

View Component
- 안드로이드 화면을 구성하는 요소들
- TextView, ImageView, Button

View Component의 종류
1. 부모(Parent)가 될 수 있는 View Component : 가지고 있느 자식 view들의 배치를 조정
  - LinearLayout
  - RelativeLayout
  - FrameLayout
  - ScrollView
  - 기타등등
  
2. 자식이 될 수 있는 Viewcomponent
  - TextView, ImageView 등,,
  
모든 ViewComponent가 가지고 있는 속성
- Width
- Height
- Background

각각의 ViewComponent들은 고유 속성으 가지고 있음
ex> TextView의 TextColor, TextSize

화면 크기 단위는 Px (화면 크기에 상관없이, 고정된 크기로 보여짐), dp로 지정 가능 (화면 크기에 비율로 반영)
=> 일반적으로 dp 사용
