window.BENCHMARK_DATA = {
  "lastUpdate": 1667632724931,
  "repoUrl": "https://github.com/yeya24/promql-engine",
  "entries": {
    "Go Benchmark": [
      {
        "commit": {
          "author": {
            "name": "yeya24",
            "username": "yeya24"
          },
          "committer": {
            "name": "yeya24",
            "username": "yeya24"
          },
          "id": "19f35c9b450de4617fc37720e1c11ed27f2a3c14",
          "message": "add continuous benchmark action",
          "timestamp": "2022-09-18T06:49:21Z",
          "url": "https://github.com/yeya24/promql-engine/pull/43/commits/19f35c9b450de4617fc37720e1c11ed27f2a3c14"
        },
        "date": 1667632724594,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRangeQuery/vector_selector",
            "value": 83750540,
            "unit": "ns/op\t28658989 B/op\t  126955 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/vector_selector",
            "value": 84061011,
            "unit": "ns/op\t28685341 B/op\t  126958 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/vector_selector",
            "value": 86823232,
            "unit": "ns/op\t28671936 B/op\t  126959 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/vector_selector",
            "value": 86745038,
            "unit": "ns/op\t28710681 B/op\t  126965 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/vector_selector",
            "value": 85844363,
            "unit": "ns/op\t28660494 B/op\t  126958 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum",
            "value": 77135020,
            "unit": "ns/op\t 9362168 B/op\t  121343 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum",
            "value": 77807316,
            "unit": "ns/op\t 9367596 B/op\t  121341 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum",
            "value": 79053008,
            "unit": "ns/op\t 9371452 B/op\t  121342 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum",
            "value": 78366679,
            "unit": "ns/op\t 9424052 B/op\t  121344 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum",
            "value": 79128876,
            "unit": "ns/op\t 9397941 B/op\t  121346 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_by_pod",
            "value": 87159385,
            "unit": "ns/op\t18829595 B/op\t  206432 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_by_pod",
            "value": 86660032,
            "unit": "ns/op\t18655584 B/op\t  206421 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_by_pod",
            "value": 86201650,
            "unit": "ns/op\t18504760 B/op\t  206401 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_by_pod",
            "value": 85058167,
            "unit": "ns/op\t18518530 B/op\t  206412 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_by_pod",
            "value": 87336119,
            "unit": "ns/op\t18602622 B/op\t  206410 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/rate",
            "value": 170956752,
            "unit": "ns/op\t30014452 B/op\t  150933 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/rate",
            "value": 172169846,
            "unit": "ns/op\t30013058 B/op\t  150932 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/rate",
            "value": 172603586,
            "unit": "ns/op\t30110693 B/op\t  150944 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/rate",
            "value": 170729373,
            "unit": "ns/op\t30095164 B/op\t  150939 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/rate",
            "value": 169768574,
            "unit": "ns/op\t30028582 B/op\t  150937 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_rate",
            "value": 166524347,
            "unit": "ns/op\t11243782 B/op\t  145361 allocs/op",
            "extra": "7 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_rate",
            "value": 168915018,
            "unit": "ns/op\t11242420 B/op\t  145351 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_rate",
            "value": 169075487,
            "unit": "ns/op\t11341824 B/op\t  145370 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_rate",
            "value": 167880334,
            "unit": "ns/op\t11268745 B/op\t  145370 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_rate",
            "value": 169223163,
            "unit": "ns/op\t11249564 B/op\t  145360 allocs/op",
            "extra": "7 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_by_rate",
            "value": 177141958,
            "unit": "ns/op\t20243448 B/op\t  230419 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_by_rate",
            "value": 173722952,
            "unit": "ns/op\t20262836 B/op\t  230434 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_by_rate",
            "value": 175222510,
            "unit": "ns/op\t20259228 B/op\t  230426 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_by_rate",
            "value": 172683602,
            "unit": "ns/op\t20239836 B/op\t  230423 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/sum_by_rate",
            "value": 172834738,
            "unit": "ns/op\t20231768 B/op\t  230411 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_one_to_one",
            "value": 49306136,
            "unit": "ns/op\t14806050 B/op\t   98602 allocs/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_one_to_one",
            "value": 47374518,
            "unit": "ns/op\t14799759 B/op\t   98598 allocs/op",
            "extra": "26 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_one_to_one",
            "value": 47659696,
            "unit": "ns/op\t14833205 B/op\t   98612 allocs/op",
            "extra": "25 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_one_to_one",
            "value": 47292915,
            "unit": "ns/op\t14809232 B/op\t   98601 allocs/op",
            "extra": "26 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_one_to_one",
            "value": 45786718,
            "unit": "ns/op\t14774314 B/op\t   98590 allocs/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_many_to_one",
            "value": 106109815,
            "unit": "ns/op\t35066171 B/op\t  192308 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_many_to_one",
            "value": 108235684,
            "unit": "ns/op\t35078873 B/op\t  192301 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_many_to_one",
            "value": 106418550,
            "unit": "ns/op\t35092832 B/op\t  192334 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_many_to_one",
            "value": 108840119,
            "unit": "ns/op\t35134450 B/op\t  192323 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_many_to_one",
            "value": 108190088,
            "unit": "ns/op\t35057427 B/op\t  192328 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_vector_and_scalar",
            "value": 91460986,
            "unit": "ns/op\t30861832 B/op\t  130909 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_vector_and_scalar",
            "value": 97403889,
            "unit": "ns/op\t30888088 B/op\t  130926 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_vector_and_scalar",
            "value": 93851488,
            "unit": "ns/op\t30942506 B/op\t  130926 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_vector_and_scalar",
            "value": 95833180,
            "unit": "ns/op\t30848889 B/op\t  130912 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/binary_operation_with_vector_and_scalar",
            "value": 95379152,
            "unit": "ns/op\t30861465 B/op\t  130913 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/unary_negation",
            "value": 97188860,
            "unit": "ns/op\t30116608 B/op\t  139064 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/unary_negation",
            "value": 92888739,
            "unit": "ns/op\t29965410 B/op\t  139045 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/unary_negation",
            "value": 90815021,
            "unit": "ns/op\t30033341 B/op\t  139056 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/unary_negation",
            "value": 93189868,
            "unit": "ns/op\t29991618 B/op\t  139046 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/unary_negation",
            "value": 91991328,
            "unit": "ns/op\t29995876 B/op\t  139047 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/vector_and_scalar_comparison",
            "value": 94673442,
            "unit": "ns/op\t30513195 B/op\t  127898 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/vector_and_scalar_comparison",
            "value": 91902033,
            "unit": "ns/op\t30527004 B/op\t  127904 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/vector_and_scalar_comparison",
            "value": 92568342,
            "unit": "ns/op\t30487254 B/op\t  127895 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/vector_and_scalar_comparison",
            "value": 91495649,
            "unit": "ns/op\t30500030 B/op\t  127900 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/vector_and_scalar_comparison",
            "value": 91899309,
            "unit": "ns/op\t30551262 B/op\t  127915 allocs/op",
            "extra": "13 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/positive_offset_vector",
            "value": 78293079,
            "unit": "ns/op\t26890844 B/op\t   98045 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/positive_offset_vector",
            "value": 79649455,
            "unit": "ns/op\t26861784 B/op\t   98044 allocs/op",
            "extra": "15 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/positive_offset_vector",
            "value": 78376483,
            "unit": "ns/op\t26852071 B/op\t   98041 allocs/op",
            "extra": "15 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/positive_offset_vector",
            "value": 79133191,
            "unit": "ns/op\t26851924 B/op\t   98040 allocs/op",
            "extra": "14 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/positive_offset_vector",
            "value": 77852423,
            "unit": "ns/op\t26855004 B/op\t   98042 allocs/op",
            "extra": "15 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/at_modifier_",
            "value": 54426799,
            "unit": "ns/op\t35151632 B/op\t   75533 allocs/op",
            "extra": "21 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/at_modifier_",
            "value": 54297743,
            "unit": "ns/op\t35151966 B/op\t   75534 allocs/op",
            "extra": "21 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/at_modifier_",
            "value": 55113216,
            "unit": "ns/op\t35151872 B/op\t   75533 allocs/op",
            "extra": "21 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/at_modifier_",
            "value": 52447763,
            "unit": "ns/op\t35151669 B/op\t   75533 allocs/op",
            "extra": "20 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/at_modifier_",
            "value": 55631610,
            "unit": "ns/op\t35151814 B/op\t   75533 allocs/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/at_modifier_with_positive_offset_vector",
            "value": 52030253,
            "unit": "ns/op\t34961725 B/op\t   69533 allocs/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/at_modifier_with_positive_offset_vector",
            "value": 53796340,
            "unit": "ns/op\t34960141 B/op\t   69534 allocs/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/at_modifier_with_positive_offset_vector",
            "value": 50609062,
            "unit": "ns/op\t34959946 B/op\t   69533 allocs/op",
            "extra": "21 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/at_modifier_with_positive_offset_vector",
            "value": 53527949,
            "unit": "ns/op\t34961940 B/op\t   69534 allocs/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/at_modifier_with_positive_offset_vector",
            "value": 51183435,
            "unit": "ns/op\t34963894 B/op\t   69533 allocs/op",
            "extra": "20 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/clamp",
            "value": 99178398,
            "unit": "ns/op\t29056739 B/op\t  130661 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/clamp",
            "value": 102295076,
            "unit": "ns/op\t29074716 B/op\t  130670 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/clamp",
            "value": 103593799,
            "unit": "ns/op\t29108371 B/op\t  130677 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/clamp",
            "value": 103568739,
            "unit": "ns/op\t29120617 B/op\t  130673 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/clamp",
            "value": 103198656,
            "unit": "ns/op\t29032737 B/op\t  130659 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/clamp_min",
            "value": 100319680,
            "unit": "ns/op\t29036400 B/op\t  130320 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/clamp_min",
            "value": 99455018,
            "unit": "ns/op\t29050939 B/op\t  130324 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/clamp_min",
            "value": 98697398,
            "unit": "ns/op\t29043180 B/op\t  130327 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/clamp_min",
            "value": 99763540,
            "unit": "ns/op\t29054915 B/op\t  130322 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/clamp_min",
            "value": 99150380,
            "unit": "ns/op\t29044147 B/op\t  130318 allocs/op",
            "extra": "12 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/complex_func_query",
            "value": 111479630,
            "unit": "ns/op\t31118253 B/op\t  135354 allocs/op",
            "extra": "9 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/complex_func_query",
            "value": 110461819,
            "unit": "ns/op\t31118548 B/op\t  135354 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/complex_func_query",
            "value": 113610926,
            "unit": "ns/op\t31120692 B/op\t  135354 allocs/op",
            "extra": "10 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/complex_func_query",
            "value": 112277276,
            "unit": "ns/op\t31139589 B/op\t  135370 allocs/op",
            "extra": "9 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/complex_func_query",
            "value": 113378103,
            "unit": "ns/op\t31173843 B/op\t  135360 allocs/op",
            "extra": "9 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/func_within_func_query",
            "value": 161762437,
            "unit": "ns/op\t30170688 B/op\t  152392 allocs/op",
            "extra": "7 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/func_within_func_query",
            "value": 159867595,
            "unit": "ns/op\t30191411 B/op\t  152402 allocs/op",
            "extra": "7 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/func_within_func_query",
            "value": 162618421,
            "unit": "ns/op\t30246980 B/op\t  152399 allocs/op",
            "extra": "7 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/func_within_func_query",
            "value": 159690501,
            "unit": "ns/op\t30312892 B/op\t  152407 allocs/op",
            "extra": "7 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/func_within_func_query",
            "value": 160567914,
            "unit": "ns/op\t30240956 B/op\t  152397 allocs/op",
            "extra": "7 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/aggr_within_func_query",
            "value": 169606988,
            "unit": "ns/op\t30179941 B/op\t  152394 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/aggr_within_func_query",
            "value": 172277644,
            "unit": "ns/op\t30272364 B/op\t  152410 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/aggr_within_func_query",
            "value": 171425104,
            "unit": "ns/op\t30197222 B/op\t  152405 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/aggr_within_func_query",
            "value": 176134876,
            "unit": "ns/op\t30501354 B/op\t  152420 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/aggr_within_func_query",
            "value": 170630427,
            "unit": "ns/op\t30188740 B/op\t  152400 allocs/op",
            "extra": "6 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/histogram_quantile",
            "value": 312948714,
            "unit": "ns/op\t96794000 B/op\t  700952 allocs/op",
            "extra": "4 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/histogram_quantile",
            "value": 309405312,
            "unit": "ns/op\t96094374 B/op\t  700934 allocs/op",
            "extra": "4 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/histogram_quantile",
            "value": 313648360,
            "unit": "ns/op\t96455304 B/op\t  700943 allocs/op",
            "extra": "4 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/histogram_quantile",
            "value": 320252831,
            "unit": "ns/op\t96498932 B/op\t  700942 allocs/op",
            "extra": "4 times\n2 procs"
          },
          {
            "name": "BenchmarkRangeQuery/histogram_quantile",
            "value": 311547879,
            "unit": "ns/op\t96499492 B/op\t  700944 allocs/op",
            "extra": "4 times\n2 procs"
          }
        ]
      }
    ]
  }
}