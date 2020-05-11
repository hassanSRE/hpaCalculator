package main

import (
	"fmt"
	"strconv"
	"math"
	"os"
	"time"

)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: hpacalc <num replicas> <utilization>")
		os.Exit(1)
	}
	replicaCount, _ := strconv.ParseInt(os.Args[1], 0, 32)
	utilization, _ := strconv.ParseInt(os.Args[2], 0, 64)

	fmt.Println("Calculating value for replicas:", replicaCount, "with AverageValue", utilization)
	statusReplicas := int32(replicaCount)
	milliValue := utilization  * int64(1000)
	_, utilizationProposal, _, _ := GetExternalPerPodMetricReplicas(statusReplicas, milliValue, "testMetric", "testNameSpace")
	fmt.Println("Decimal Value: ", utilizationProposal)

}

// GetExternalMetricReplicas calculates the desired replica count based on a
// target metric value (as a milli-value) for the external metric in the given
// namespace, and the current replica count.
func GetExternalPerPodMetricReplicas(statusReplicas int32, targetUtilizationPerPod int64, metricName, namespace string) (replicaCount int32, utilization int64, timestamp time.Time, err error) {
	utilization = targetUtilizationPerPod
	tolerance := float64(0)
	replicaCount = statusReplicas
	usageRatio := float64(utilization) / (float64(targetUtilizationPerPod) * float64(replicaCount))
	if math.Abs(1.0-usageRatio) > tolerance {
		// update number of replicas if the change is large enough
		replicaCount = int32(math.Ceil(float64(utilization) / float64(targetUtilizationPerPod)))
	}
	utilization = int64(math.Ceil(float64(utilization) / float64(statusReplicas)))

	return replicaCount, utilization, timestamp, nil
}
